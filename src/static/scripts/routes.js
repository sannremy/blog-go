import MainPage from './pages/main';
import StoryPage from './pages/story';
import PortfolioPage from './pages/portfolio';

export default class Routes {
  /**
   * Routes contructor
   */
  constructor () {
    this.pages = {
      main: new MainPage(),
      story: new StoryPage(),
      portfolio: new PortfolioPage()
    };

    this.activePageKey = null;
    this.activePage = null;
  }

  /**
   * Index handler
   * @return {void}
   */
  indexHandler = () => {
    if (this.activePageKey !== 'main') {
      if (this.activePage) {
        this.activePage.unload();
      }

      this.pages.main.load();

      this.activePageKey = 'main';
      this.activePage = this.pages[this.activePageKey];
    }
  }

  /**
   * Story handler
   * @return {void}
   */
  storyHandler = () => {
    if (this.activePageKey !== 'story') {
      if (this.activePage) {
        this.activePage.unload();
      }

      this.pages.story.load();

      this.activePageKey = 'story';
      this.activePage = this.pages[this.activePageKey];
    }
  }

  /**
   * Portfolio handler
   * @return {void}
   */
  portfolioHandler = () => {
    if (this.activePageKey !== 'portfolio') {
      if (this.activePage) {
        this.activePage.unload();
      }

      this.pages.portfolio.load();

      this.activePageKey = 'portfolio';
      this.activePage = this.pages[this.activePageKey];
    }
  }
}
