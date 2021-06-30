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

export const CREATE_ARTICLE_MUTATION = gql`
  mutation createArticle {
    createArticle(input: {author: "next front", title: "from next front desu", content: "hello world from next front"}) {
      id
      author
      title
      content
    }
  }
`

// mutation createArticle {
//   createArticle(input: {author:"line", title: "line title", content: "line content hello"}) {
//     id
//     author
//     title
//     content
//   }
// }

export interface ArticlesData {
  articles: Article[];
}

export interface ArticleData {
  article: Article;
}