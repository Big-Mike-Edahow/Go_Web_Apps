Simple Go CRUD web app. It uses SQLite3 for the database. HTML and CSS. No frameworks.

One thing in particular that was difficult for me to figure out was how to use 𝗿𝗮𝗻𝗴𝗲 to loop over data in the HTML Template. A slice[ ] of structs{ } is sent to the template, and then {{𝗿𝗮𝗻𝗴𝗲 .}}...{{end}} is used to access the struct.

In the case where only one struct is sent to the template, the fields of the struct are accessed with {{.Title}}, {{.Author}}, etc.

Hopefully this helps someone out.
