import MainPage from './pages/main';
import StoryPage from './pages/story';
import PortfolioPage from './pages/portfolio';

export default class Routes {
  /**
   * Routes contructor
   */
  constructor() {
    this.pages = {
      main: new MainPage(),
      story: new StoryPage(),
      portfolio: new PortfolioPage()
    };
  }

  /**
   * Index handler
   * @return {void}
   */
  indexHandler = () => {
    this.pages.main.load();
  }

  /**
   * Story handler
   * @return {void}
   */
  storyHandler = () => {
    this.pages.story.load();
  }

  /**
   * Portfolio handler
   * @return {void}
   */
  portfolioHandler = () => {
    this.pages.portfolio.load();
  }
}
