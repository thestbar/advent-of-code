require_relative '../utils'

# Part 1
grid = Utils.build_grid('input.txt', ' ')
results = []

Utils.read_cells(grid) do |cell, r, c|
  next if r == grid.length - 1

  if grid[grid.length - 1][c] == '+'
    results[c] ||= 0
    results[c] += cell.strip.to_i
  else
    results[c] ||= 1
    results[c] *= cell.strip.to_i
  end
end

puts "Part 1: #{results.sum}"

# Part 2
grid = Utils.build_grid('input.txt')
results = []

m = grid.length
n = grid[0].length

operator = '+'
op_index = -1
for c in 0...n
  # Get the operator
  curr_op = grid[m - 1][c].strip
  unless curr_op.empty?
    operator = curr_op
    op_index += 1
  end

  # Build number as a string
  str_num = ''
  for r in 0...m - 1
    str_num += grid[r][c].strip
  end

  next if str_num.empty?

  # Convert to integer and apply operation
  num = str_num.to_i

  if operator == '+'
    results[op_index] ||= 0
    results[op_index] += num
  else
    results[op_index] ||= 1
    results[op_index] *= num
  end
end

puts "Part 2: #{results.sum}"
