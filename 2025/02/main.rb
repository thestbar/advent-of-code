# Part 1
File.open('input.txt', 'r') do |file|
  invalid_ids_sum = 0
  file.each_line do |line|
    ranges = line.strip.split(',')
    ranges.each do |range|
      start_str, end_str = range.split('-')
      start_num = start_str.to_i
      end_num = end_str.to_i
      puts "Start: #{start_num}, End: #{end_num}"
      (start_num..end_num).each do |num|
        str_num = num.to_s
        len = str_num.length
        next if len.odd?

        half = len / 2
        first_half = str_num[0...half]
        second_half = str_num[half...len]
        next unless first_half == second_half

        puts "Found invalid ID: #{str_num}"
        invalid_ids_sum += num
      end
    end
  end

  puts "Total sum of invalid IDs: #{invalid_ids_sum}"
end

# Part 2
File.open('input.txt', 'r') do |file|
  invalid_ids_sum = 0
  file.each_line do |line|
    ranges = line.strip.split(',')
    ranges.each do |range|
      start_str, end_str = range.split('-')
      start_num = start_str.to_i
      end_num = end_str.to_i
      puts "Start: #{start_num}, End: #{end_num}"
      (start_num..end_num).each do |num|
        str_num = num.to_s
        len = str_num.length
        chunk_size = len / 2
        while chunk_size >= 1
          puts "Checking number: #{str_num} with chunk size #{chunk_size}"
          # Create an array of all the chunks of size chunk_size
          # e.g. for 123123 and chunk_size 3, chunks = [123, 123]
          # for 123123123 and chunk_size 3, chunks = [123, 123, 123]
          str_num = num.to_s
          len = str_num.length
          if len % chunk_size != 0
            chunk_size -= 1
            next
          end

          chunks = str_num.scan(/.{1,#{chunk_size}}/)
          if chunks.uniq.length == 1
            puts "Found invalid ID: #{str_num} with chunk size #{chunk_size}"
            invalid_ids_sum += num
            break
          end
          chunk_size -= 1
        end
      end
    end
  end

  puts "Total sum of invalid IDs: #{invalid_ids_sum}"
end
