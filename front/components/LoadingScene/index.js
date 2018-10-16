import React from "react";
import styled, { keyframes } from "styled-components";

const Container = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  padding-top: 120px;
`;

const animation = keyframes`
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
`;

const Loading = styled.div`
  border-top: 12px solid rgba(255, 255, 255, 0.5);
  border-right: 12px solid rgba(255, 255, 255, 0.5);
  border-bottom: 12px solid rgba(255, 255, 255, 0.5);
  border-left: 12px solid #ffffff;
  transform: translateZ(0);
  animation: ${animation} 1.1s infinite linear;
  &,
  &:after {
    border-radius: 50%;
    width: 120px;
    height: 120px;
  }
`;

export default function LoadingScene() {
  return (
    <Container>
      <Loading />
    </Container>
  );
}