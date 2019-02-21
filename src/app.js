const express = require('express');
const app = express();
const port = process.env.PORT || '8080';

// Template engine
app.set('views', './src/views')
app.set('view engine', 'pug');

// Serve static
app.use(express.static('dist'));

// i18n
const languagesAvailable = ['en', 'fr'];
const acceptLanguage = require('accept-language');
acceptLanguage.languages(languagesAvailable);

// Load language modules
const languages = {};
for (let i = 0; i < languagesAvailable.length; i++) {
  languages[languagesAvailable[i]] = require(`./i18n/${languagesAvailable[i]}`);
}

// Index route
app.get(`/:language(${languagesAvailable.join('|')})?`, (req, res) => {
  // Get client language
  let clientLanguage = req.headers['accept-language'] || 'en';

  // Override client language if set in param
  if (req.params.language) {
    clientLanguage = req.params.language;
  }

  // Get available language
  const language = acceptLanguage.get(clientLanguage);

  res.render('index', {
    language: language,
    __: (key) => {
      return languages[language][key] || key;
    }
  });
});

// HTTP Server
app.listen(port, () => console.log(`App listening on port ${port}!`));
