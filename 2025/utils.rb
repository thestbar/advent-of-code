module Utils
  DIRS = [[0, 1], [1, 1], [1, 0], [1, -1], [0, -1], [-1, -1], [-1, 0], [-1, 1]].freeze

  def self.build_grid(path, split_char = '')
    grid = []
    File.open(path, 'r') do |file|
      file.each_line do |line|
        row = line.chomp.split(split_char)
        grid << row
      end
    end

    grid
  end

  def self.print_grid(grid)
    grid.each do |row|
      puts row.join(' ')
    end
  end

  def self.read_cells(grid, &block)
    grid.each_with_index do |row, r|
      row.each_with_index do |cell, c|
        block.call(cell, r, c)
      end
    end
  end

  def self.within_bounds?(grid, row, col)
    row >= 0 && row < grid.length && col >= 0 && col < grid[0].length
  end

  def self.get_neighbors(grid, row, col)
    neighbors = []
    DIRS.each do |dr, dc|
      new_row = row + dr
      new_col = col + dc
      neighbors << [new_row, new_col] if within_bounds?(grid, new_row, new_col)
    end

    neighbors
  end
end
