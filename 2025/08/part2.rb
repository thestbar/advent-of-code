require_relative '../utils'
require_relative 'circuits'

module Part2
  def self.run
    grid = Utils.build_grid('input.txt', ',')
    circuits = []
    junctions = grid.map do |row|
      junction = Junction.new(row[0].to_i, row[1].to_i, row[2].to_i)
      circuits << junction.circuit

      junction
    end

    slds = []
    junctions.combination(2) do |p1, p2|
      slds << { pair: [p1, p2], distance: p1.sld(p2) }
    end

    slds.sort_by! { |entry| entry[:distance] }
    last_pair = nil
    res = nil

    slds.each do |entry|
      break res = last_pair if circuits.size == 1

      j1, j2 = entry[:pair]
      last_pair = entry

      next if j1.circuit == j2.circuit

      circuit1 = j1.circuit
      circuit2 = j2.circuit

      circuit2.junctions.each do |junction|
        circuit1.add_juction!(junction)
      end

      circuits.delete(circuit2)
    end

    # Output results
    puts "Part 2: #{res[:pair][0].x * res[:pair][1].x}"
  end
end
