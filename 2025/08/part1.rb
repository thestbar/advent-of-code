require_relative '../utils'
require_relative 'circuits'

module Part1
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

    count = 0
    limit = 1000
    slds.each do |entry|
      break if count >= limit

      j1, j2 = entry[:pair]

      if j1.circuit != j2.circuit
        circuit1 = j1.circuit
        circuit2 = j2.circuit

        circuit2.junctions.each do |junction|
          circuit1.add_juction!(junction)
        end

        circuits.delete(circuit2)
      end

      count += 1
    end

    # Sort circuits by number of junctions descending
    circuits.sort_by! { |circuit| -circuit.junctions.size }

    res_p1 = circuits.first(3).reduce(1) do |prod, circuit|
      prod * circuit.junctions.size
    end

    # Output results
    puts "Part 1: #{res_p1}"
  end
end
