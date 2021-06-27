import { gql } from '@apollo/client';
import { Article } from '../../types';

export const ARTICLES_QUERY = gql`
  query ListArticle {
    articles {
      id
      author
      title
      content
    }
  }
`;

export interface ArticlesData {
  articles: Article[];
}