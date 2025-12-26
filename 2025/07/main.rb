require_relative '../utils'

# Part 1
def move_beam(grid, position)
  new_position = [position[0] + 1, position[1]]
  return 0 unless Utils.within_bounds?(grid, new_position[0], new_position[1])
  return 0 if grid[new_position[0]][new_position[1]] == '|'

  cell = grid[new_position[0]][new_position[1]]
  if cell == '^'
    left_position = [new_position[0], new_position[1] - 1]
    right_position = [new_position[0], new_position[1] + 1]
    grid[left_position[0]][left_position[1]] = '|'
    grid[right_position[0]][right_position[1]] = '|'
    left_beam = move_beam(grid, left_position)
    right_beam = move_beam(grid, right_position)

    return 1 + left_beam + right_beam
  end

  grid[new_position[0]][new_position[1]] = '|'
  move_beam(grid, new_position)
end

grid = Utils.build_grid('input.txt')
start = Utils.find_all_occurrences(grid, 'S').first
puts "Part 1: #{move_beam(grid, start)}"

# Part 2
def count_worlds(grid, row, col, memo)
  return 1 if row == grid.length && col >= 0 && col < grid[0].length
  return 0 unless Utils.within_bounds?(grid, row, col)

  cell = grid[row][col]
  key = "#{row},#{col}"
  return memo[key] if memo.key?(key)

  if cell == '^'
    left_worlds = count_worlds(grid, row, col - 1, memo)
    right_worlds = count_worlds(grid, row, col + 1, memo)
    memo[key] = left_worlds + right_worlds
    return memo[key]
  end

  count_worlds(grid, row + 1, col, memo)
end

puts "Part 2: #{count_worlds(grid, start[0], start[1], {})}"
