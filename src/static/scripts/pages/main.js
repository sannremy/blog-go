import Page from './_page';
import TypeIt from 'typeit';

const hello = document.querySelector('#hello');
const storyButtons = document.querySelector('#see-story');
const menu = document.querySelector('#menu');

export default class MainPage extends Page {
  isDisplayedOnce = false;
  container = document.querySelector('#main-page');

  /**
   * Constructor
   */
  constructor () {
    super();
  }

  /**
   * Load a page
   * @return {void}
   */
  load = () => {
    if (this.isDisplayedOnce) {
      this.container.classList.remove('hide');
    } else {
      this.isDisplayedOnce = true;

      new TypeIt('#hello', {
        speed: 50,
        nextStringDelay: 100,
        cursorChar: '<span class="cursor">|</span>',
        waitUntilVisible: true,
        afterComplete: (instance) => {
          // Hide blinking cursor
          document.querySelector('.ti-cursor').classList.add('hide');

          // Show buttons
          storyButtons.classList.remove('hide');
          storyButtons.classList.add('animated', 'faster', 'fadeInUp');

          // Show menu
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
  }

  /**
   * Unload a page
   * @return {void}
   */
  unload = () => {
    this.container.classList.add('hide');
  }
}
