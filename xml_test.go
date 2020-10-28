package replace

import (
	`errors`
	`io/ioutil`
	`os`
	`testing`
)

const (
	androidXMLFilename = ".testdata/AndroidManifest.xml"
	androidXML         = `
<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
          android:versionName="0.4.0"
          package="com.class330.tv">
    
    <extra>
        <username>测试用户</username>
    </extra>

    <application
            android:allowBackup="true"
            android:icon="@mipmap/ic_launcher"
            android:label="@string/app_name"
            android:supportsRtl="true"
            android:theme="@style/AppTheme">
        <activity android:name=".MainActivity">
            <intent-filter>
                <action android:name="android.intent.action.MAIN"/>

                <category android:name="android.intent.category.LAUNCHER"/>
            </intent-filter>
        </activity>
    </application>

</manifest>
`
	expectedAndroidXML = `
<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
          android:versionName="0.4.0"
          package="com.class330.tv">
    
    <extra>
        <username>Storezhang</username>
    </extra>

    <application
            android:allowBackup="true"
            android:icon="@mipmap/ic_launcher"
            android:label="测试应用"
            android:supportsRtl="true"
            android:theme="@style/AppTheme">
        <activity android:name=".MainActivity">
            <intent-filter>
                <action android:name="android.intent.action.MAIN"/>

                <category android:name="android.intent.category.LAUNCHER"/>
            </intent-filter>
        </activity>
    </application>

</manifest>
`
)

func TestXML(t *testing.T) {
	if err := ioutil.WriteFile(androidXMLFilename, []byte(androidXML), os.ModePerm); nil != err {
		t.Error(err)
	}

	json := NewXMLReplace(androidXMLFilename, XMLReplaceElement{
		XPath: "/manifest",
		Type:  XMLReplaceTypeAttr,
		Key:   "android:versionName",
		Value: "1.0.0",
	}, XMLReplaceElement{
		XPath: "/manifest/application",
		Type:  XMLReplaceTypeAttr,
		Key:   "android:label",
		Value: "测试应用",
	}, XMLReplaceElement{
		XPath: "/manifest/extra",
		Type:  XMLReplaceTypeText,
		Key:   "username",
		Value: "Storezhang",
	})

	if err := json.Replace("./"); nil != err {
		t.Error(err)
	}

	if fileData, err := ioutil.ReadFile(androidXMLFilename); nil != err {
		t.Error(err)
	} else if expectedAndroidXML != string(fileData) {
		t.Error(errors.New("文件修改失败"))
	}
}
