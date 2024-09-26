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
    "shinjuku_b1": false,
    "shinjuku_b2": false,
    "sakurajosui_b1": false,
    "sakurajosui_b2": false,
    "sakurajosui_b3": false,
    "sakurajosui_b4": false,
    "sakurajosui_b5": false,
    "sakurajosui_b6": false,
    "chofu_b1": false,
    "chofu_b2": false,
    "chofu_b3": false,
    "chofu_b4": false,
    "chofu_b5": false,
    "hashimoto_b1": false,
    "hashimoto_b2": false,
    "hachioji_b1": false,
    "hachioji_b2": false,
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
        <svg width="100%" viewBox="0 0 540 200">
          <rect width={540} height={200} fill="#222222" />
          {/* 新宿 */}
          <Platform position={{ x: 480, y: 110 }} />
          <Platform position={{ x: 480, y: 170 }} />
          {/* 桜上水 */}
          <Platform position={{ x: 360, y: 110 }} />
          <Platform position={{ x: 360, y: 170 }} />
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
            id="chofu_b5"
            points={[
              { x: 260, y: 120 },
              { x: 320, y: 120 },
            ]}
          />
          <Section
            id="sakurajosui_b3"
            points={[
              { x: 320, y: 120 },
              { x: 400, y: 120 },
            ]}
          />
          <Section
            id="sakurajosui_b4"
            points={[
              { x: 320, y: 120 },
              { x: 340, y: 100 },
              { x: 380, y: 100 },
              { x: 400, y: 120 },
            ]}
          />
          <Section
            id="sakurajosui_b6"
            points={[
              { x: 400, y: 120 },
              { x: 500, y: 120 },
            ]}
          />
          <Section
            id="shinjuku_b2"
            points={[
              { x: 500, y: 120 },
              { x: 500, y: 160 },
              { x: 460, y: 160 },
            ]}
          />
          <Section
            id="shinjuku_b1"
            points={[
              { x: 460, y: 160 },
              { x: 400, y: 160 },
            ]}
          />
          <Section
            id="sakurajosui_b1"
            points={[
              { x: 400, y: 160 },
              { x: 380, y: 180 },
              { x: 340, y: 180 },
              { x: 320, y: 160 },
            ]}
          />
          <Section
            id="sakurajosui_b2"
            points={[
              { x: 400, y: 160 },
              { x: 320, y: 160 },
            ]}
          />
          <Section
            id="sakurajosui_b5"
            points={[
              { x: 320, y: 160 },
              { x: 260, y: 160 },
            ]}
          />
          {/* 八王子方面 */}
          <Section
            id="chofu_b2"
            points={[
              { x: 260, y: 160 },
              { x: 240, y: 140 },
              { x: 200, y: 140 },
            ]}
          />
          <Section
            id="chofu_b4"
            points={[
              { x: 200, y: 140 },
              { x: 140, y: 80 },
              { x: 40, y: 80 },
            ]}
          />
          <Section
            id="hachioji_b1"
            points={[
              { x: 40, y: 80 },
              { x: 40, y: 40 },
              { x: 80, y: 40 },
            ]}
          />
          <Section
            id="hachioji_b2"
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
            id="chofu_b1"
            points={[
              { x: 260, y: 160 },
              { x: 200, y: 160 },
            ]}
          />
          <Section
            id="chofu_b3"
            points={[
              { x: 200, y: 160 },
              { x: 40, y: 160 },
            ]}
          />
          <Section
            id="hashimoto_b1"
            points={[
              { x: 40, y: 160 },
              { x: 40, y: 120 },
              { x: 80, y: 120 },
            ]}
          />
          <Section
            id="hashimoto_b2"
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
