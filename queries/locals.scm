;; Define locals for ini grammar
(
  ; Define the scope for sections using the section_name node
  (section
    (section_name) @scope)

  ; Define the setting_name as a variable definition within a section
  (setting
    (setting_name) @definition.var)
)
