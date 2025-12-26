collecting_ingredients = true
ranges = []
in_range_count = 0

# Part 1: Read ranges and check numbers
File.open('input.txt', 'r') do |file|
  file.each_line do |line|
    line = line.strip
    next collecting_ingredients = false if line.empty?

    if collecting_ingredients
      puts "Collecting ingredients: #{line}"
      data = line.split('-').map(&:to_i)
      start = data[0]
      ending = data[1]
      ranges << (start..ending)
    else
      puts "Checking number: #{line}"
      number = line.to_i
      ranges.each do |range|
        if range.include?(number)
          in_range_count += 1
          break
        end
      end
    end
  end
end

puts "Part 1 (Total numbers in range): #{in_range_count}"

# Part 2: Check for full containment
ranges.sort_by! { |r| [r.begin, r.end] }

puts "Sorted Ranges: #{ranges.inspect}"

# Merge overlapping ranges
merged_ranges = []
ranges.each do |current_range|
  if merged_ranges.empty? || merged_ranges.last.end < current_range.begin
    merged_ranges << current_range
  else
    last_range = merged_ranges.pop
    new_range = (last_range.begin..[last_range.end, current_range.end].max)
    merged_ranges << new_range
  end
end

puts "Merged Ranges: #{merged_ranges.inspect}"

fully_contained_count = merged_ranges.sum do |range|
  range.end - range.begin + 1
end

puts "Part 2 (Total fully contained numbers): #{fully_contained_count}"
