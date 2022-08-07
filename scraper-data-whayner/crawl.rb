
require 'webdrivers'
require 'json'
require 'open-uri'
require 'net/http'

regDesired = 140
screen_height = 9999
startTime = Time.now
print "Start ", startTime, "\n"

# chromium
options = Selenium::WebDriver::Options.chrome
options.add_argument("--headless")
options.add_argument("--incognito")
driver = Selenium::WebDriver.for :chrome, options: options
wait = Selenium::WebDriver::Wait.new(:timeout => 10)

# file
filePath = './data.json'
topicsFilePath = './topics.txt'
root_url_list = Array.new

# read topics 
File.readlines(topicsFilePath).each do |line|
  
  # line not empty
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
  print 'Scaning Root: ', root_url

  # vars
  regObtained = 0
  regPrevious = regObtained
  sameRegCount = 0

  # TODO ADD PRESENCE WAIT
  # driver.manage.timeouts.implicit_wait = 1000
  driver.get(root_url)

  # get only the desired amount

  while (regObtained < regDesired) do

    # get video entries

    videoE_list = driver.find_elements(css: 'ytd-video-renderer')
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
  for i in 0..(regObtained-1) do

    videoE = videoE_list[i]
    titleEl = videoE.find_element(css: 'a#video-title')

    # get attributes
    title = titleEl.attribute('title')
    url = titleEl.attribute('href')

    # description, tags and thumbnail
    html = Net::HTTP.get_response(URI.parse(url)).body
    description = html.scan(/"shortDescription":"(?<!\\)[\s\S]+?"/)[0][20..-2]
    tags = html.scan(/\<meta name="keywords" content="(?<!\\)[\s\S]+?"/)[0][31..-2]
    thumbnail = html.scan(/thumbnailUrl" href="(?<!\\)[\s\S]+?"/)[0][20..-2]

    # save
    item = {
      'url' => url,
      'title' => title,
      'description' => description,
      'tags' => tags,
      'thumbnail' => thumbnail
    } 

    # puts item
    print (i +1), ' ', url

    # write item to file
    File.open(filePath, "a") do |line|
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

# time stats
endTime = Time.now
print "\nStart ", startTime, "\n"
print "Finish ", endTime, "\n"
print "Total ", (endTime - startTime), "\n"
print " sec or ", (endTime - startTime)/60, " min\n"

# exit
driver.quit

