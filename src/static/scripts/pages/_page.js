export default class Page {
  /**
   * Load a page
   * @return {void}
   */
  load () {
    throw new Error('Page.load has to be override');
  }

  /**
   * Unload a page
   * @return {void}
   */
  unload () {
    throw new Error('Page.unload has to be override');
  }
}
