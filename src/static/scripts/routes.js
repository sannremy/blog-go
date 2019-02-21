import TypeIt from 'typeit';

export default class Routes {
  /**
   * Index handler
   * @return {void}
   */
  indexHandler () {
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
        expandable.classList.add('big-h');

        setTimeout(() => {
          expandable.classList.add('big-v');
        }, 500);
      }, 500);
    });
  }

  /**
   * Story handler
   * @return {void}
   */
  storyHandler () {
    console.log('yes');
  }
}
