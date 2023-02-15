import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import './Post.css'


  function Post(props) {
    // const [text, setText] = useState('');
    console.log(props);
    return (
      <div className='post-container'>
        <div className='post-header'>
        <p>{props.date}</p>
        <p>{props.title}</p>
        <p>{props.author}</p>
        </div>
        
        <div className='post-body'>
            <p>{props.content}</p>
        </div>
        <Link to={`/post/${props.postid}`}>
        <button>View Post</button>
        </Link>
      </div>
    );
  }
export default Post;