#
# Simple script to test getting websites
# Just like curl
#

require 'open-uri'
require 'net/http'
require 'json'

url = "https://www.youtube.com/watch?v=-kW8OEclIWQ"

# uri = URI.parse(url)
# response = Net::HTTP.get_response(uri)
# html = response.body

html = Net::HTTP.get_response(URI.parse(url)).body

desc = html.scan(/"shortDescription":"(?<!\\)[\s\S]+?"/)[0][20..-2]
tags = html.scan(/\<meta name="keywords" content="(?<!\\)[\s\S]+?"/)[0][31..-2]
thumbnail = html.scan(/thumbnailUrl" href="(?<!\\)[\s\S]+?"/)[0][20..-2]

item = {
  'desc' => desc,
  'thumbnail' => thumbnail,
  'tags' => tags
}

puts item

# puts JSON.parse('{' + keywords + '}') 
# puts exa["shortDescription"]
