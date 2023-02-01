import {StyleSheet, SafeAreaView} from 'react-native';
import { FC, ReactNode } from 'react'

type SafeAreaContainerProps = {
  children: ReactNode
}

export const SafeAreaContainer: FC<SafeAreaContainerProps> = ({children}) => {
  return (
    <SafeAreaView style={styles.container}>
      {children}
    </SafeAreaView>
  )
}
const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  text: {
    fontSize: 25,
    fontWeight: '500',
  },
});
