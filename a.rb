require "net/http"
require "json"

uri = URI.parse("https://api.sampleapis.com/coffee/hot")
res = Net::HTTP.get_response(uri)
items = JSON.parse(res.body, symbolize_names: true)

items.each {|item|
   p item[:title]
   p item[:description]
   p '=========================='
}
