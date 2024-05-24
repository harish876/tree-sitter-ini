module.exports = grammar({
    name: 'ini',
  
    extras: $ => [
      $.comment,
      $._blank,
      /[\t ]/
    ],
  
    rules: {
      document: $ => seq(
        repeat($._blank),  // Eat blank lines at top of file.
        repeat($.section),
      ),
  
      // Section has:
      // - a title
      // - zero or more settings (name=value pairs)
      section: $ => seq(
        $.section_name,
        repeat($.setting),
      ),
  
      section_name: $ => seq(
        '[',
        alias(/[^\[\]\n]+/, $.text),
        ']',
        '\n',
      ),
  
      setting: $ => seq(
        alias(/[^;#=\s\[]+/, $.setting_name),
        '=',
        alias(/[^;#\n]+/, $.setting_value),
        '\n',
      ),
  
      comment: $ => seq(/[;#]/, alias(/.*/, $.text), '\n'),
  
      _blank: () => field('blank', '\n'),
    }
  });