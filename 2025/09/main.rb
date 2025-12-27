require_relative '../utils'

class Point
  attr_accessor :x, :y

  def initialize(x, y)
    @x = x
    @y = y
  end

  def sld(other)
    ((other.x - @x)**2 + (other.y - @y)**2)**0.5
  end

  def area(other)
    ((@x - other.x).abs + 1) * ((@y - other.y).abs + 1)
  end
end

points = []
File.open('input.txt', 'r') do |file|
  file.each_line do |line|
    vals = line.chomp.split(',')
    row = vals[1].to_i
    col = vals[0].to_i
    points << Point.new(row, col)
  end
end

puts points.inspect

slds = []
points.combination(2) do |p1, p2|
  slds << { pair: [p1, p2], distance: p1.sld(p2), area: p1.area(p2) }
end

slds.sort_by! { |entry| -entry[:area] }

puts slds

puts "Largest Rectangle Area: #{slds.first[:area]}"
