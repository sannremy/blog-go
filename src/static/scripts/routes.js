import MainPage from './pages/main';
import StoryPage from './pages/story';
import PortfolioPage from './pages/portfolio';

export default class Routes {
  /**
   * Index handler
   * @return {void}
   */
  indexHandler () {
    const mainPage = new MainPage();
    mainPage.load();
  }

  /**
   * Story handler
   * @return {void}
   */
  storyHandler () {
    const storyPage = new StoryPage();
    storyPage.load();
  }

  /**
   * Portfolio handler
   * @return {void}
   */
  portfolioHandler () {
    const portfolioPage = new PortfolioPage();
    portfolioPage.load();
  }
}
