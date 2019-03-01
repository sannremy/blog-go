import '../styles/main.scss';

import Routes from './routes';
import page from 'page';
import 'pace-progress';
import FontFaceObserver from 'fontfaceobserver';

const fontA = new FontFaceObserver('Work Sans');
const languages = ['en', 'fr'];
const routes = new Routes();

fontA.load().then(() => {
  const regexLang = languages.join('|');
  page(`/:language(${regexLang})?`, routes.indexHandler);
  page(`/:language(${regexLang})?/story`, routes.storyHandler);
  page(`/:language(${regexLang})?/portfolio`, routes.portfolioHandler);
  page();
});
