import React from 'react';
import { useParams } from 'react-router-dom';

function PostPage() {
  const { id } = useParams();

  // fetch post data using the ID

  return (
    <div>
      <h1>Post {id}</h1>
      {/* display post data */}
    </div>
  );
}

export default PostPage;