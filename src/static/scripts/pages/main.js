import Page from './_page';
import TypeIt from 'typeit';

const hello = document.querySelector('#hello');
const storyButtons = document.querySelector('#see-story');

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
      cursorChar: '<span class="cursor">|</span>',
      waitUntilVisible: true,
      afterComplete: (instance) => {
        storyButtons.classList.remove('hide');
        storyButtons.classList.add('animated', 'faster', 'fadeInUp');
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
  }

  /**
   * Unload a page
   * @return {void}
   */
  unload () {

  }
}
