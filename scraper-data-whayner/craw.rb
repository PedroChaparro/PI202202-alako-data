
require 'webdrivers'
require 'json'

# chromium
options = Selenium::WebDriver::Options.chrome
options.add_argument("--headless")
driver = Selenium::WebDriver.for :chrome, options: options
wait = Selenium::WebDriver::Wait.new(:timeout => 10)

filePath = './data.json'
topicsFilePath = './topics.txt'
root_url_list = Array.new

# read topics from file

File.readlines(topicsFilePath).each do |line|
  
  # line is not empty
  if (!/\A[[:space:]]*\z/.match?(line))
    root_url_list.append("https://www.youtube.com/results?search_query=" + line.gsub(' ', '+'))
  end

end

puts 'URLS: '
puts root_url_list

# add json opening

File.open(filePath, "w") do |line|
  line.puts '{ "videos": ['
end

# search each root url

for k in 0..(root_url_list.length-1) do

  root_url = root_url_list[k]
  print 'Scaning: ', root_url

  # vars
  url_list = Array.new
  screen_height = 9999
  regDesired = 3
  regObtained = 0
  regPrevious = regObtained
  sameRegCount = 0

  # TODO ADD PRESENCE WAIT
  driver.get(root_url)
  driver.manage.timeouts.implicit_wait = 1000

  # get only the desired amount

  while (regObtained < regDesired) do

    # get video entries

    videoE_list = driver.find_elements(css: 'ytd-video-renderer a#video-title')
    regPrevious = regObtained
    regObtained = videoE_list.length

    if (regPrevious = regObtained) 
      sameRegCount += 1
    end

    if (sameRegCount >= 20) 
      break
    end

    # scroll down
    driver.action
      .scroll_by(0, screen_height)
      .perform

  end

  # crop reg quantity
  regObtained = [regObtained, regDesired].min

  # add video url to list
  videoE_list.each do |i|

    # respect item limit
    if url_list.length < regObtained
      url_list.append(i.attribute('href'))
    else
      break
    end

  end

  # go through every video
  
  for i in 0..(regObtained-1) do

    iurl = url_list[i]
    print i, ' ', iurl

    # get
    driver.get(iurl)

    # wait till title element is accesible
    
    titleE = wait.until {driver.find_element(css: '#info h1 yt-formatted-string')}
    tagsE = driver.find_element({name: 'keywords'})
    thumbnailE = driver.find_elements(css: 'body > div#watch7-content > link')

    # extract values
    
    tags = tagsE.attribute('content')
    title = titleE.attribute('innerText')
    thumbnail = ''

    # extract thumbnail
    
    thumbnailE.each do |thumb|
      if thumb.attribute('itemprop') == 'thumbnailUrl'
        thumbnail = thumb.attribute('href')
      end
    end

    # print debug info
    # print "\nurl: ", iurl
    # print "\ntitle: ", title
    # print "\ntags: ", tags
    # print "\nthumbnail: ", thumbnail
    # print "\n"

    item = {
      'url' => iurl,
      'title' => title,
      'tags' => tags,
      'thumbnail' => thumbnail
    }

    # write item to file
    
    File.open(filePath,"a") do |line|
      toPut = JSON[item]

      # last item
      
      if ((i != regObtained -1) || (k != (root_url_list.length) -1))
        toPut += ','
      end

      line.puts toPut # write
      puts "\033[0;32m ✔\033[0m" # print with colors just because
    end

  end
end

# add json closing

File.open(filePath,"a") do |line|
  line.puts ']}'
end

# exit

driver.quit

