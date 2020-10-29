package replace

import (
	`encoding/xml`
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
)

const (
	expectedVersionName   = "1.0.0"
	expectedAndroidLabel  = "测试应用"
	expectedExtraUsername = "Storezhang"
)

type AndroidManifest struct {
	XMLName xml.Name `xml:"android manifest"`

	VersionName string `xml:"android versionName,attr"`

	Extra struct {
		Username string `xml:"username"`
	} `xml:"extra"`

	Application struct {
		AndroidLabel string `xml:"android label,attr"`
	} `xml:"application"`
}

func TestXML(t *testing.T) {
	if err := ioutil.WriteFile(androidXMLFilename, []byte(androidXML), os.ModePerm); nil != err {
		t.Error(err)
	}

	replace := NewXMLReplace(androidXMLFilename, XMLReplaceElement{
		XPath: "/manifest",
		Type:  XMLReplaceTypeAttr,
		Key:   "android:versionName",
		Value: expectedVersionName,
	}, XMLReplaceElement{
		XPath: "/manifest/application",
		Type:  XMLReplaceTypeAttr,
		Key:   "android:label",
		Value: expectedAndroidLabel,
	}, XMLReplaceElement{
		XPath: "/manifest/extra",
		Type:  XMLReplaceTypeText,
		Key:   "username",
		Value: expectedExtraUsername,
	})

	if err := replace.Replace("./"); nil != err {
		t.Error(err)
	}

	if fileData, err := ioutil.ReadFile(androidXMLFilename); nil != err {
		t.Error(err)
	} else {
		manifest := &AndroidManifest{}
		if err = xml.Unmarshal(fileData, manifest); nil != err {
			t.Error(errors.New("文件修改失败"))
		}

		if expectedVersionName != manifest.VersionName || expectedAndroidLabel != manifest.Application.AndroidLabel || expectedExtraUsername != manifest.Extra.Username {
			t.Error(errors.New("文件修改失败"))
		}
	}
}
