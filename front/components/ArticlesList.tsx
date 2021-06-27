// import Link from 'next/link'
import {useQuery} from '@apollo/client'
import { ARTICLES_QUERY, ArticlesData} from '../graphql/queries/articles.query'
import { NextPage } from 'next'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface ArticlesListProps {}

const ArticlesList: NextPage<ArticlesListProps> = () => {
  const { loading, error, data } = useQuery<ArticlesData>(ARTICLES_QUERY);
  if (loading) return <p>now loading...</p>;
  if (error) return <p>Error desu: { JSON.stringify(error) }</p>;
  const { articles } = data;
  if (!articles) return <p>nothing articles data</p>;
  // eslint-disable-next-line no-console
  console.log(articles);
  return (
    <ul>
      <li>list</li>
    </ul>
  )
}

export default ArticlesList;