import {useQuery} from '@apollo/client'
import { NextPage } from 'next'
import { ArticlesData, ARTICLES_QUERY } from '../../graphql/queries/articles.query'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface IndexProps {}

const Index: NextPage<IndexProps> = () => {
  // const { data } = useQuery<ArticlesData>(ARTICLES_QUERY);
  // const { articles } = data
  // if (!articles) return <p>nothing to data</p> 
  return (
    <>
      <h2>hello world list</h2>
    </>
  )
}

export default Index