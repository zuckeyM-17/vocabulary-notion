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

def append_database_row(word_en, word_ja)
  client = Notion::Client.new

  database_id = ENV['DATABASE_ID']
  properties = {
    'English' => { title: [{ text: { content: word_en } }] },
    'Japanese' => { rich_text: [{ text: { content: word_ja } }] }
  }
  parent = { database_id: }
  client.create_page(parent:, properties:)
end

def main
  setup
  word_en = ARGV[0]
  word_ja = translate(word_en)
  append_database_row(word_en, word_ja)
  puts "#{word_en}: #{word_ja}"
end

main
