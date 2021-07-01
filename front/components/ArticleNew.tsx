import { useMutation } from "@apollo/client";
import { NextPage } from "next";
import { useState } from "react";
import { CREATE_ARTICLE_MUTATION} from '../graphql/queries/articles.query'
import ArticlesList from './ArticlesList'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface ArticleNewProps {}

const ArticleNew: NextPage<ArticleNewProps> = () => {
  const [author, setAuthor] = useState("")
  const [title, setTitle] = useState("")
  const [content, setContent] = useState("")
  const [createArticle] = useMutation(CREATE_ARTICLE_MUTATION, {
    variables: {
      author: author,
      title: title,
      content: content
    }
  });

  return (
    <>
      <h1>hello article new</h1>
      <input type="text" value={author} placeholder="author" onChange={e => setAuthor(e.target.value)} />
      <br />
      <input type="text" value={title} placeholder="title" onChange={e => setTitle(e.target.value)} />
      <br />
      <input type="text" value={content} placeholder="content" onChange={e => setContent(e.target.value)} />
      <br />
      <button onClick={() => createArticle()}>mutation execution</button>
      <ArticlesList />
    </>
  )
}

export default ArticleNew