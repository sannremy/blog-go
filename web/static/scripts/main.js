import '../styles/main.scss';

// Prism: code highlighter
import Prism from 'prismjs/components/prism-core';
import 'prismjs/components/prism-clike';
import 'prismjs/components/prism-bash';
import 'prismjs/components/prism-http';
import 'prismjs/components/prism-markup-templating';
import 'prismjs/components/prism-php';
import 'prismjs/components/prism-javascript';
import 'prismjs/components/prism-yaml';
import 'prismjs/themes/prism.css';

const postElement = document.querySelector('.post');
if (postElement) {
  Prism.highlightAllUnder(postElement);
}
