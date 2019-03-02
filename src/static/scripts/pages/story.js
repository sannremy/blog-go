import Page from './_page';

export default class StoryPage extends Page {
  container = document.querySelector('#story-page');

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
    console.log('/story');
    this.container.classList.remove('hide');
  }

  /**
   * Unload a page
   * @return {void}
   */
  unload = () => {
    this.container.classList.add('hide');
  }
}
