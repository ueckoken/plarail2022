import type { NextPage } from "next";
import Head from "next/head";
import { useEffect, useState } from "react";
import { io } from "socket.io-client";
import Platform from "../components/Platform";
import Section from "../components/Section";
import { BlocksContext } from "../contexts";
import { validateSignalTaikenMessage } from "../lib/signal";
import styles from "../styles/Home.module.css";
import { Blocks } from "../types";

const Home: NextPage = () => {
  const [blocks, setBlocks] = useState<Blocks>({
    s0: false,
    s1: false,
    s2: false,
    s3: false,
    s4: false,
    s5: false,
    s6: false,
    s7: false,
    s8: false,
    s9: false,
    s10: false,
    s11: false,
    s12: false,
  });

  useEffect(() => {
    if (process.env.NEXT_PUBLIC_NO_WEBSOCKET) {
      const sequence: string[][] = [
        ["s0"],
        ["s1"],
        ["s2"],
        ["s3"],
        ["s4"],
        ["s5"],
        ["s6"],
        ["s11"],
        ["s12"],
        ["s0"],
        ["s1"],
        ["s2"],
        ["s7"],
        ["s8"],
        ["s9"],
        ["s10"],
        ["s11"],
        ["s12"],
      ];

      let index = sequence.length - 1;

      const interval = setInterval(() => {
        index = (index + 1) % sequence.length;
        const blocks: Blocks = {};
        for (const section of sequence[index]) {
          blocks[section] = true;
        }
        setBlocks(blocks);
      }, 1000);

      return () => {
        clearInterval(interval);
      };
    } else {
      // before_first_request を起動
      fetch("http://localhost:50050");

      const socket = io("http://localhost:50050");
      socket.on("signal_taiken", (data: unknown) => {
        console.log(data);
        if (validateSignalTaikenMessage(data)) {
          setBlocks(data.blocks);
        } else {
          throw new Error(`Illegal message: ${JSON.stringify(data)}`);
        }
      });

      return () => {
        socket.close();
      };
    }
  }, []);

  return (
    <div>
      <Head>
        <title>自動運転ビジュアライザ</title>
      </Head>
      <BlocksContext.Provider value={blocks}>
        <svg width="100%" viewBox="0 0 480 200">
          <rect width={480} height={200} fill="#222222" />
          {/* 新宿 */}
          <Platform position={{ x: 420, y: 110 }} />
          <Platform position={{ x: 420, y: 170 }} />
          {/* 桜上水 */}
          <Platform position={{ x: 340, y: 110 }} />
          <Platform position={{ x: 340, y: 170 }} />
          {/* 調布 */}
          <Platform position={{ x: 220, y: 110 }} />
          <Platform position={{ x: 220, y: 150 }} />
          {/* 八王子 */}
          <Platform position={{ x: 60, y: 30 }} />
          <Platform position={{ x: 60, y: 90 }} />
          {/* 橋本 */}
          <Platform position={{ x: 60, y: 110 }} />
          <Platform position={{ x: 60, y: 170 }} />

          {/* 新宿方面 */}
          <Section
            id="s11"
            points={[
              { x: 260, y: 120 },
              { x: 360, y: 120 },
            ]}
          />
          <Section
            id="s12"
            points={[
              { x: 360, y: 120 },
              { x: 440, y: 120 },
            ]}
          />
          <Section
            id="s0"
            points={[
              { x: 440, y: 120 },
              { x: 440, y: 160 },
              { x: 400, y: 160 },
            ]}
          />
          <Section
            id="s1"
            points={[
              { x: 400, y: 160 },
              { x: 320, y: 160 },
            ]}
          />
          <Section
            id="s2"
            points={[
              { x: 320, y: 160 },
              { x: 260, y: 160 },
            ]}
          />
          {/* 八王子方面 */}
          <Section
            id="s3"
            points={[
              { x: 260, y: 160 },
              { x: 240, y: 140 },
              { x: 200, y: 140 },
            ]}
          />
          <Section
            id="s4"
            points={[
              { x: 200, y: 140 },
              { x: 140, y: 80 },
              { x: 40, y: 80 },
            ]}
          />
          <Section
            id="s5"
            points={[
              { x: 40, y: 80 },
              { x: 40, y: 40 },
              { x: 80, y: 40 },
            ]}
          />
          <Section
            id="s6"
            points={[
              { x: 80, y: 40 },
              { x: 140, y: 40 },
              { x: 200, y: 100 },
              { x: 240, y: 100 },
              { x: 260, y: 120 },
            ]}
          />
          {/* 橋本方面 */}
          <Section
            id="s7"
            points={[
              { x: 260, y: 160 },
              { x: 200, y: 160 },
            ]}
          />
          <Section
            id="s8"
            points={[
              { x: 200, y: 160 },
              { x: 40, y: 160 },
            ]}
          />
          <Section
            id="s9"
            points={[
              { x: 40, y: 160 },
              { x: 40, y: 120 },
              { x: 80, y: 120 },
            ]}
          />
          <Section
            id="s10"
            points={[
              { x: 80, y: 120 },
              { x: 260, y: 120 },
            ]}
          />
        </svg>
      </BlocksContext.Provider>
    </div>
  );
};

export default Home;
