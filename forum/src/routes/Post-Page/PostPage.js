import React from 'react';
import { useParams } from 'react-router-dom';
import Header from '../../components/Header/Header';
import Body from '../../components/Body/Body';
function PostPage() {
  const { id } = useParams();

  // fetch post data using the ID

  return (
    <>
      <Header />
      <Body />
      <div>
        <h1>Post {id}</h1>
        {/* display post data */}
      </div>
    </>
  );
}

export default PostPage;