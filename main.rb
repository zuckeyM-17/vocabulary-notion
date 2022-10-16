# frozen_string_literal: true

require 'dotenv'
require 'deepl'
require 'notion_ruby_client'

Dotenv.load

def setup
  DeepL.configure do |config|
    config.auth_key = ENV['DEEPL_API_KEY']
    config.host = 'https://api-free.deepl.com'
  end

  Notion.configure do |config|
    config.token = ENV['NOTION_TOKEN']
  end
end

def translate(english_word)
  translation = DeepL.translate english_word, nil, 'JA'
  translation.text
end

def append_database_row(word)
  client = Notion::Client.new

  database_id = ENV['DATABASE_ID']
  properties = {
    'English' => { title: [{ text: { content: word } }] },
    'Japanese' => { rich_text: [{ text: { content: translate(word) } }] }
  }
  parent = { database_id: }
  client.create_page(parent:, properties:)
end

def main
  setup
  append_database_row(ARGV[0])
  puts 'success!'
end

main
