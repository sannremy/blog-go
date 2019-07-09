import '../styles/main.scss';

// Prism: code highlighter
import Prism from 'prismjs/components/prism-core';
import 'prismjs/components/prism-clike';
import 'prismjs/components/prism-bash';
import 'prismjs/themes/prism.css';

const postElement = document.querySelector('.post');
if (postElement) {
  Prism.highlightAllUnder(postElement);
}
