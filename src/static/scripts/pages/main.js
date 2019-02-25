import Page from './_page';
import TypeIt from 'typeit';

const hello = document.querySelector('#hello');
const storyButtons = document.querySelector('#see-story');
const menu = document.querySelector('#menu');

export default class MainPage extends Page {
  /**
   * Load a page
   * @return {void}
   */
  load () {
    // Show page
    document.body.classList.remove('opacity-0');

    new TypeIt('#hello', {
      speed: 50,
      nextStringDelay: 100,
      cursorChar: '<span class="cursor">|</span>',
      waitUntilVisible: true,
      afterComplete: (instance) => {
        storyButtons.classList.remove('hide');
        storyButtons.classList.add('animated', 'faster', 'fadeInUp');

        menu.classList.remove('hide');
        menu.classList.add('animated', 'faster', 'slideInDown');
      },
      strings: [
        `Hello!`,
        `My name is <span class="bold">Sann-Remy</span>.`,
        `I'm a <span class="bold">Software Engineer</span>. :)`
      ]
    })
    .pause(500)
    .delete(1)
    .pause(100)
    .delete(1)
    .pause(250)
    .type('\ud83d\ude0a')
    .pause(100)
    .go();
  }

  /**
   * Unload a page
   * @return {void}
   */
  unload () {

  }
}
