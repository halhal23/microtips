import Link from 'next/link'
import ArticlesList from '../components/ArticlesList';
const Test = (): JSX.Element => {
  return (
    <>
      <div>test test</div>
      <Link href="/">top page</Link>
      <ArticlesList />
    </>
  )
}

export default Test;