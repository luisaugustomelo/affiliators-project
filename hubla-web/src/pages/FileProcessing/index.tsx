import { useState, useEffect } from 'react';
import axios from 'axios';
import { StyledTable, StyledTableCell } from './styles';

interface DataItem {
  Balance: number;
  Role: number;
  Name: string;
}

interface Data {
  id: string;
}
function FileProcessing({ id }: Data) {
  const [balances, setBalance] = useState<DataItem[]>([]);

  useEffect(() => {
    const fetchData = async (intervalId: number) => {
      try {
        const response = await axios.get(`/checkStatus/${id}`);
        setBalance(response.data);
        console.log(response.data);
        if (response.data) {
          clearInterval(intervalId);
        }
      } catch (error) {
        console.error('Error:', error);
      }
    };

    const intervalId = window.setInterval(() => fetchData(intervalId), 5000);

    return () => {
      clearInterval(intervalId);
    };
  }, [id]);

  return (
    <div>
      <StyledTable>
        <thead>
          <tr>
            <StyledTableCell>NOME</StyledTableCell>
            <StyledTableCell>VALOR</StyledTableCell>
            <StyledTableCell>PAPEL</StyledTableCell>
          </tr>
        </thead>
        <tbody>
          {balances != null &&
            balances.map((item: DataItem, index) => (
              <tr key={index}>
                <StyledTableCell>{item.Name}</StyledTableCell>
                <StyledTableCell>R$ {item.Balance.toFixed(2)}</StyledTableCell>
                <StyledTableCell>
                  {item.Role == 1 ? 'creator' : 'affiliated'}
                </StyledTableCell>
              </tr>
            ))}
        </tbody>
      </StyledTable>
    </div>
  );
}

export default FileProcessing;
