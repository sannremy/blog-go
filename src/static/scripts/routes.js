import TypeIt from 'typeit';

const hello = document.querySelector('#hello');
const storyExpandable = document.querySelector('#story-expendable');
const storyButtons = document.querySelector('#see-story');
const storyButton = document.querySelector('#story-button');

export default class Routes {
  /**
   * Index handler
   * @return {void}
   */
  indexHandler () {
    // Show page
    document.body.classList.remove('opacity-0');

    const storyButtonRect = storyButton.getBoundingClientRect();
    storyExpandable.style.top = storyButtonRect.top + 'px';
    storyExpandable.style.left = storyButtonRect.left + 'px';
    storyExpandable.style.width = storyButtonRect.width + 'px';
    storyExpandable.style.height = storyButtonRect.height + 'px';

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

    storyButton.addEventListener('click', () => {
      hello.classList.add('animated', 'faster', 'fadeOut');

      setTimeout(() => {
        storyExpandable.classList.add('big-h');

        setTimeout(() => {
          storyExpandable.classList.add('big-v');
        }, 500);
      }, 500);
    });
  }

  /**
   * Story handler
   * @return {void}
   */
  storyHandler () {
  }
}
