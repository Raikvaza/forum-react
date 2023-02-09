import React, { useState } from 'react';
import './Post.css'


  function Post(props) {
    // const [text, setText] = useState('');
    return (
      <div className='post-container'>
        <div className='post-header'>
        <p>Time</p>
        <p>Title</p>
        <p>Author</p>
        </div>
        
        <div className='post-body'>
            <p>Some Text</p>
        </div>

      </div>
    );
  }
export default Post;