import React, { useCallback, Dispatch, SetStateAction } from 'react';
import { useDropzone } from 'react-dropzone';
import axios from 'axios';

interface DataProps {
  setID: Dispatch<SetStateAction<string>>;
}

const FileUpload: React.FC<DataProps> = ({ setID }: DataProps) => {
  const onDrop = useCallback(
    (file: File[]) => {
      if (file[0].type !== 'text/plain') {
        alert(
          'This file type is not allowed. Only .txt files can be uploaded.',
        );
        return;
      }

      const formData = new FormData();
      formData.append('file', file[0]);
      const token = localStorage.getItem('@Hubla:token');
      // const user = localStorage.getItem('@Hubla:user');
      axios
        .post('/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: token,
          },
        })
        .then(response => {
          console.log(response.data);
          setID(response.data.ID);
        })
        .catch(error => {
          console.error('Error:', error);
          setID('0');
        });
    },
    [setID],
  );

  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

  return (
    <div
      {...getRootProps()}
      style={{
        background: '#D7FF60',
        border: '2px dashed gray',
        padding: '20px',
        marginTop: '20px',
        marginLeft: '20%',
        marginRight: '20%',
      }}
    >
      <input {...getInputProps()} />
      {isDragActive ? (
        <p>Drop the files here...</p>
      ) : (
        <p>Drag in drop some file (.txt) here, or click to select file</p>
      )}
    </div>
  );
};

export default FileUpload;
