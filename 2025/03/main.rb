def find_max_joltage(bank, digits = 2)
  max_value = 0
  curr_max_index = 0
  curr_max = bank[curr_max_index].to_i

  while digits > 0
    init_curr_max_index = curr_max_index
    init_curr_max = curr_max

    curr_max_index, curr_max = find_max_index(curr_max_index, curr_max, bank)

    while curr_max_index >= bank.length - digits + 1
      curr_max_index, curr_max = find_max_index(init_curr_max_index, init_curr_max, bank, bank.length - digits + 1)
      init_curr_max_index = curr_max_index
      init_curr_max = curr_max
    end

    max_value += curr_max * (10**(digits - 1))

    curr_max_index += 1
    curr_max = bank[curr_max_index].to_i

    digits -= 1
  end

  max_value
end

def find_max_index(curr_max_index, curr_max, bank, length = bank.length)
  for i in curr_max_index...length
    curr = bank[i].to_i
    if curr > curr_max
      curr_max = curr
      curr_max_index = i
    end
  end

  [curr_max_index, curr_max]
end

File.open('input.txt', 'r') do |file|
  part1_res = 0
  part2_res = 0
  file.each_line do |line|
    line = line.strip
    part1_res += find_max_joltage(line)
    part2_res += find_max_joltage(line, 12)
  end

  puts "Final Result Part1: #{part1_res}"
  puts "Final Result Part2: #{part2_res}"
end
