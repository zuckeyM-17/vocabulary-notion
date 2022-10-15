# frozen_string_literal: true

require 'dotenv'
require 'deepl'

Dotenv.load

def translate(english_word)
  DeepL.configure do |config|
    config.auth_key = ENV['DEEPL_API_KEY']
    config.host = 'https://api-free.deepl.com'
  end

  translation = DeepL.translate english_word, nil, 'JA'
  translation.text
end

puts translate('dog')
puts translate('cat')

