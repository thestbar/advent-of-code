require_relative '../utils'

grid = Utils.build_grid('input.txt')
Utils.print_grid(grid)

# Part 1
res = 0
Utils.read_cells(grid) do |cell, r, c|
  next unless cell == '@'

  neighbors = Utils.get_neighbors(grid, r, c)
  num_of_rolls = neighbors.select { |nr, nc| grid[nr][nc] == '@' }.count
  res += 1 if num_of_rolls < 4
end

puts "Part 1 Result: #{res}"

# Part 2
res = 0
while true
  remove = []

  Utils.read_cells(grid) do |cell, r, c|
    next unless cell == '@'

    neighbors = Utils.get_neighbors(grid, r, c)
    num_of_rolls = neighbors.select { |nr, nc| grid[nr][nc] == '@' }.count
    if num_of_rolls < 4
      res += 1
      remove << [r, c]
    end
  end

  break if remove.empty?

  remove.each do |r, c|
    grid[r][c] = '.'
  end
end

puts "Part 2 Result: #{res}"
