import TypeIt from 'typeit';

export default class Routes {
  container = document.querySelector('#main-page');

  /**
   * Routes contructor
   */
  constructor () {
  }

  /**
   * Index handler
   * @return {void}
   */
  indexHandler = () => {
    new TypeIt('#about-intro', {
      speed: 50,
      nextStringDelay: 100,
      cursorChar: '<span class="cursor">|</span>',
      waitUntilVisible: true,
      afterComplete: (instance) => {
        // Hide blinking cursor
        document.querySelector('.ti-cursor').classList.add('hide');

        // Show buttons
        document.querySelector('#about-button').classList.remove('hide');
        document.querySelector('#about-button').classList.add(
          'animated',
          'fadeInUp'
        );
      },
      strings: [
        `Hello!`,
        `My name is <span class="upper-bold">Sann-Remy</span>.`,
        `I'm a <span class="upper-bold">Software Engineer</span>. :)`
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
