import Page from './_page';

export default class PortfolioPage extends Page {
  container = document.querySelector('#portfolio-page');

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
    console.log('/portfolio');
  }

  /**
   * Unload a page
   * @return {void}
   */
  unload = () => {
    this.container.classList.add('opacity-0');
  }
}
