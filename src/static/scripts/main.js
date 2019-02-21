import animateCSS from 'animate.css';
import '../styles/main.scss';

import TypeIt from 'typeit';
import FontFaceObserver from 'fontfaceobserver';

var fontA = new FontFaceObserver('Work Sans');

fontA.load().then(() => {
  // Show page
  document.body.classList.remove('opacity-0');

  const hello = document.querySelector('#hello');
  const expandable = document.querySelector('#expandable');
  const storyButtons = document.querySelector('#see-story');
  const storyButton = document.querySelector('#story-button');

  const storyButtonRect = storyButton.getBoundingClientRect();
  expandable.style.top = storyButtonRect.top + 'px';
  expandable.style.left = storyButtonRect.left + 'px';
  expandable.style.width = storyButtonRect.width + 'px';
  expandable.style.height = storyButtonRect.height + 'px';
  expandable.innerText = storyButton.textContent;

  new TypeIt('#hello', {
    speed: 50,
    cursorChar: '<span class="cursor">|</span>',
    waitUntilVisible: true,
    afterComplete: (instance) => {
      storyButtons.classList.add('animated', 'faster', 'fadeInUp');
      storyButtons.classList.remove('opacity-0');
    },
    strings: [
      `Hello!`,
      `My name is <span class="bold">Sann-Remy</span>.`,
      `I'm a <span class="bold">Software Engineer</span>. :)`
    ]
  })
  .pause(1000)
  .delete(1)
  .pause(250)
  .delete(1)
  .pause(500)
  .type('\ud83d\ude0a')
  .pause(500)
  .go();

  storyButton.addEventListener('click', () => {
    hello.classList.add('animated', 'faster', 'fadeOut');

    setTimeout(() => {
      expandable.classList.add('big-h');

      setTimeout(() => {
        expandable.classList.add('big-v');
      }, 500);
    }, 500);
  });
});
