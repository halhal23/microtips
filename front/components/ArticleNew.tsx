import { useMutation } from "@apollo/client";
import { NextPage } from "next";
import { CREATE_ARTICLE_MUTATION} from '../graphql/queries/articles.query'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface ArticleNewProps {}

const ArticleNew: NextPage<ArticleNewProps> = () => {
  const [createArticle] = useMutation(CREATE_ARTICLE_MUTATION, {});

  return (
    <>
      <h1>hello article new</h1>
      <button onClick={() => createArticle()}>mutation execution</button>
    </>
  )
}

export default ArticleNew