def rotate(pos, input)
  rotation = input[0]
  steps = input[1..].to_i

  sum =
    if rotation == 'R'
      pos + steps
    else
      pos - steps
    end

  [sum % 100, (sum / 100).abs]
end

File.open('input.txt', 'r') do |file|
  pos = 50
  zero_count = 0
  sum_zeros_from_rotation = 0
  file.each_line do |line|
    result, zeros_from_rotation = rotate(pos, line.strip)
    puts "Starting at #{pos}, after instruction #{line.strip}, "\
         "new position is #{result}, passed over zero #{zeros_from_rotation} times"
    zero_count += 1 if result.zero?
    sum_zeros_from_rotation += zeros_from_rotation
    pos = result
  end
  puts "Final position: #{pos}"

  # Part 1
  puts "Password (Number of times position was zero): #{zero_count}"

  # Part 2
  puts "Enhanced Password (Sum of zeros from rotations): #{sum_zeros_from_rotation - 1}"
end

# Brute force solution for finding a specific password
File.open('input.txt', 'r') do |file|
  pos = 50
  zeros = 0
  file.each_line do |line|
    line = line.strip
    rotation = line[0]
    steps = line[1..].to_i

    while steps.positive?
      if rotation == 'R'
        pos += 1
      else
        pos -= 1
      end

      pos = 0 if pos == 100
      pos = 99 if pos == -1

      zeros += 1 if pos.zero?

      steps -= 1
    end
  end

  puts "Brute Force Final Position: #{pos}, Zeros Count: #{zeros}"
end
