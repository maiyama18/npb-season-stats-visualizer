import React from 'react';
import { LineSerieData, LineDatum, ResponsiveLine } from '@nivo/line';
import { connect } from 'react-redux';
import { AppState } from '../../store/reducers';
import { BatterStats, BatterStatType, PitcherStats, PitcherStatType } from '../../store/types';

interface StateProps {
  graphData: LineSerieData[];
  useMesh: boolean;
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
    xScale={{
      type: 'time',
      format: '%Y-%m-%d',
      precision: 'day',
    }}
    xFormat={'time:%Y-%m-%d'}
    axisBottom={{
      format: '%m/%d',
      tickValues: 'every 2 days',
    }}
    yScale={{
      type: 'linear',
      min: 'auto',
      max: 'auto',
    }}
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
    animate={true}
    useMesh={props.useMesh}
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
    let data: LineDatum[];
    switch (playersType) {
      case 'pitchers':
        const pStat = (p.stats as PitcherStats)[graphStat as PitcherStatType];
        data = pStat.dates.map((dateStr, i) => ({
          x: dateStr,
          y: pStat.values[i],
        }));
        break;
      case 'batters':
        const bStat = (p.stats as BatterStats)[graphStat as BatterStatType];
        data = bStat.dates.map((dateStr, i) => ({
          x: dateStr,
          y: bStat.values[i],
        }));
        break;
      default:
        data = [];
    }

    return {
      id: p.name,
      data,
    };
  });

  const useMesh = canUseMesh(graphData);

  return {
    graphData,
    useMesh,
  };
};

const canUseMesh = (graphData: LineSerieData[]): boolean => {
  let dataPoints: LineDatum[] = [];
  graphData.forEach(d => {
    dataPoints = [...dataPoints, ...d.data];
  });

  if (dataPoints.length < 3) {
    return false;
  }

  let xs = new Set();
  let ys = new Set();
  dataPoints.forEach(p => {
    xs.add(p.x);
    ys.add(p.y);
  });

  return xs.size > 1 && ys.size > 1;
};

export default connect(
  mapStateToProps,
  {}
)(StatsGraph);
