require 'erb'
require 'optparse'

# Define default variable values
defaults = {
  image_name_1: 'marcin0/minus-na-plus',
  image_name_2: 'marcin0/minus-na-plus',
  image_name_3: 'marcin0/minus-na-plus',
  image_tag_1: 'proxy',
  image_tag_2: 'fastify',
  image_tag_3: 'fiber',
  environment: 'production',
  database_host: 'db',
  memory_limit: '512M',
  target: 'dev'
}

OptionParser.new do |opts|
    opts.banner = 'Usage: generate-docker-compose.rb [options]'

    opts.on('-n1', '--name1 NAME', 'Docker image name') do |name|
        defaults[:image_name_1] = name
    end

    opts.on('-n2', '--name2 NAME', 'Docker image name') do |name|
        defaults[:image_name_2] = name
    end

    opts.on('-n3', '--name3 NAME', 'Docker image name') do |name|
        defaults[:image_name_3] = name
    end

    opts.on('-t1', '--tag1 TAG1', 'Docker image tag1') do |tag|
        defaults[:image_tag_1] = tag
    end

    opts.on('-t2', '--tag2 TAG2', 'Docker image tag2') do |tag|
        defaults[:image_tag_2] = tag
    end

    opts.on('-t3', '--tag3 TAG3', 'Docker image tag3') do |tag|
        defaults[:image_tag_3] = tag
    end

    opts.on('-e', '--environment ENV', 'Application environment') do |env|
        defaults[:environment] = env
    end

    opts.on('-d', '--database HOST', 'Database host') do |host|
        defaults[:database_host] = host
    end

    opts.on('-m', '--memory LIMIT', 'Container memory limit') do |limit|
        defaults[:memory_limit] = limit
    end
end.parse!

template_path = File.join(__dir__, 'compose.yml.erb')
template = File.read(template_path)

context = Object.new
defaults.each do |key, value|
  context.instance_variable_set("@#{key}", value)
end

result = ERB.new(template).result(context.instance_eval { binding })

puts result