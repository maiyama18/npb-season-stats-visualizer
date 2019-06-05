import React from 'react';
import { LineSerieData, LineDatum, ResponsiveLine } from '@nivo/line';
import { connect } from 'react-redux';
import { AppState } from '../../store/reducers';
import { BatterStats, BatterStatType, PitcherStats, PitcherStatType } from '../../store/types';

interface StateProps {
  graphData: LineSerieData[];
}

const StatsGraph = (props: StateProps) => (
  <ResponsiveLine
    data={props.graphData}
    margin={{
      top: 20,
      bottom: 70,
      left: 70,
      right: 110,
    }}
    // xScale={{
    //   type: 'time',
    //   format: 'native',
    // }}
    xScale={{
      type: 'linear',
      min: 'auto',
      max: 'auto',
    }}
    yScale={{
      type: 'linear',
      min: 'auto',
      max: 'auto',
    }}
    animate={true}
    legends={[
      {
        anchor: 'bottom-right',
        direction: 'column',
        translateX: 100,
        justify: false,
        itemWidth: 80,
        itemHeight: 20,
        symbolSize: 12,
        symbolShape: 'circle',
      },
    ]}
    theme={{
      background: '#f4f4f4',
      axis: {
        legend: {
          text: {
            fontSize: 20,
          },
        },
      },
    }}
  />
);

const mapStateToProps = (state: AppState): StateProps => {
  const { graphStat, selectedPlayers, playersType } = state.players;

  const graphData: LineSerieData[] = selectedPlayers.map(p => {
    console.log(p);
    let data: LineDatum[];
    switch (playersType) {
      case 'pitchers':
        const pStat = (p.stats as PitcherStats)[graphStat as PitcherStatType];
        console.log(pStat);
        data = pStat.dates.map((dateStr, i) => {
          const d = new Date(dateStr);
          console.log(d, d.getTime());
          return {
            x: new Date(d).getTime(),
            y: pStat.values[i],
          };
        });
        break;
      case 'batters':
        const bStat = (p.stats as BatterStats)[graphStat as BatterStatType];
        console.log(bStat);
        data = bStat.dates.map((dateStr, i) => {
          const d = new Date(dateStr);
          console.log(d, d.getTime());
          return {
            x: new Date(d).getTime(),
            y: bStat.values[i],
          };
        });
        break;
      default:
        data = [];
    }

    return {
      id: p.name,
      data,
    };
  });

  return {
    graphData,
  };
};

export default connect(
  mapStateToProps,
  {}
)(StatsGraph);
