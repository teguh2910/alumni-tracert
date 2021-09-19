import React, {useState, useEffect} from 'react';
import {
  SafeAreaView,
  ScrollView,
  StatusBar,
  StyleSheet,
  Image,
  View,
} from 'react-native';
import {useDispatch, useSelector} from 'react-redux';
import {Button, Text, Card, Chip} from 'react-native-elements';
import {PricingCard} from 'react-native-elements';
import ImagePicker from 'react-native-image-crop-picker';

const DetailCertificate = ({navigation}) => {
  const {isLogin, detailCertificate} = useSelector(state => ({
    isLogin: state.isLogin,
    detailCertificate: state.detailCertificate,
  }));
  const [isLoading, setLoading] = useState(false);
  const [isLoadingIjazah, setLoadingIjazah] = useState(false);
  const [isLoadingTranscript, setLoadingTranscript] = useState(false);
  const onLegalisir = () => {};
  const onUploadIjazah = () => {
    ImagePicker.openCamera({
      width: 300,
      height: 400,
      cropping: true,
    }).then(image => {
      console.log(image);
    });
  };
  const onUploadTranscript = () => {};
  const labelStyle = {
    marginTop: 12,
    fontWeight: 'bold',
    fontSize: 16,
    color: '#9CA3AF',
  };
  return (
    <SafeAreaView style={{backgroundColor: '#ffffff', flex: 1}}>
      <ScrollView
        contentContainerStyle={{
          padding: 24,
          backgroundColor: '#ffffff',
        }}>
        <Text style={{fontSize: 28, fontWeight: 'bold'}}>
          {detailCertificate.majorStudy}
        </Text>
        <View style={{marginTop: 24}}>
          <Text style={labelStyle}>NIM</Text>
          <Text style={{fontSize: 20, fontWeight: 'bold'}}>
            {detailCertificate.nim}
          </Text>
          <Text style={labelStyle}>Lulus</Text>
          <Text style={{fontSize: 20, fontWeight: 'bold'}}>
            {detailCertificate.graduationYear}
          </Text>
          <Text style={labelStyle}>No. Ijazah</Text>
          <Text style={{fontSize: 20, fontWeight: 'bold'}}>
            {detailCertificate.noIjazah}
          </Text>
        </View>
        {detailCertificate.legalize.status === 0 && (
          <View style={{flex: 1, paddingTop: 24}}>
            <Card
              containerStyle={{
                margin: 0,
                borderStyle: 'dashed',
                borderRadius: 1,
              }}>
              <Card.Title>Upload Ijazah</Card.Title>
              <Card.Divider />
              <Button
                title="Upload"
                buttonStyle={{
                  backgroundColor: '#047857',
                }}
                loading={isLoadingIjazah}
                onPress={onUploadIjazah}
              />
            </Card>
            <Card
              containerStyle={{
                margin: 0,
                marginTop: 24,
                borderStyle: 'dashed',
                borderRadius: 1,
              }}>
              <Card.Title>Upload Transkrip</Card.Title>
              <Card.Divider />
              <Button
                title="Upload"
                buttonStyle={{
                  backgroundColor: '#047857',
                }}
                loading={isLoadingTranscript}
                onPress={onUploadTranscript}
              />
            </Card>
          </View>
        )}
      </ScrollView>
      <View style={{padding: 24}}>
        <Button
          title="Legalisir"
          buttonStyle={{
            backgroundColor: '#047857',
          }}
          loading={isLoading}
          onPress={onLegalisir}
        />
      </View>
    </SafeAreaView>
  );
};

export default DetailCertificate;
