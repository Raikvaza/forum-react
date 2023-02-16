import React from 'react';
import { useParams } from 'react-router-dom';
import Layout from '../../components/Layout/Layout';
import Post from '../../components/Post/Post';
function PostPage() {
  const { id } = useParams();



  // fetch post data using the ID

  return (
    <>
    <Layout>
      <Post id={id}/>
      {/* Set it as params in Post component */}
    </Layout>
    </>
  );
}

export default PostPage;